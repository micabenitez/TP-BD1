package funciones

import (
	db "alpuin-benitez-chianese-zanardi-tp/lib"
	"fmt"
)

func CreateFuncAndTrigger() {
	FuncionAutorizarCompra()
	FuncionAlertaRechazo()
	TriggerAlertaRechazo()
	FuncionGenerarResumen()
	FuncionGenerarConsumos()
	FuncionAlertaCompra()
	TriggerAlertaCompra()
	FuncionAlertaLimite()
	TriggerAlertaLimite()
}

func DeleteFuncAndTrigger() {
	db.ExecuteQuery("drop function if exists autorizarCompra cascade;")
	db.ExecuteQuery("drop trigger if exists alerta_rechazo_trg on rechazo cascade;")
	db.ExecuteQuery("drop function if exists alerta_rechazo cascade;")
	db.ExecuteQuery("drop function if exists generarResumen cascade;")
	db.ExecuteQuery("drop function if exists generarConsumos cascade;")
	db.ExecuteQuery("drop function if exists alerta_compra cascade;")
	db.ExecuteQuery("drop trigger if exists alerta_compra_trg on compra cascade;")
	db.ExecuteQuery("drop function if exists alerta_limite cascade;")
	db.ExecuteQuery("drop trigger if exists alerta_limite_trg on rechazo cascade;")
}

func FuncionGenerarConsumos() {
	db.ExecuteQuery(`create or replace function generarConsumos() returns void as $$ 
    declare
        c record;
    begin
        for c in select * from consumo
            loop                
                PERFORM autorizarCompra(c.nrotarjeta,c.codseguridad,c.nrocomercio,c.monto);                                
            end loop;            
    end;
    $$language plpgsql;`)
}

func GenerarConsumos() {
	db.ExecuteQuery(`select generarConsumos();`)
}

/** autorización de compra se deberá incluir la lógica que reciba los datos de una
compra—número de tarjeta, código de seguridad, número de comercio y monto—y
que devuelva true si se autoriza la compra ó false si se rechaza. El procedimiento
deberá validar los siguientes elementos antes de autorizar:
– Que el número de tarjeta sea existente, y que corresponda a alguna tarjeta vigente.
En caso de que no cumpla, se debe cargar un rechazo con el mensaje ?tarjeta no
válida ó no vigente.
– Que el código de seguridad sea el correcto. En caso de que no cumpla, se debe
cargar un rechazo con el mensaje ?código de seguridad inválido.
– Que el monto total de compras pendientes de pago más la compra a realizar no
supere el límite de compra de la tarjeta. En caso de que no cumpla, se debe cargar
un rechazo con el mensaje ?supera límite de tarjeta.
– Que la tarjeta no se encuentre vencida. En caso de que no cumpla, se debe cargar
un rechazo con el mensaje ?plazo de vigencia expirado.
– Que la tarjeta no se encuentre suspendida. En caso que no cumpla, se debe cargar
un rechazo con el mensaje la tarjeta se encuentra suspendida.
Si se aprueba la compra, se deberá guardar una fila en la tabla compra, con los datos
de la compra.
*/

func FuncionAutorizarCompra() {
	funcSQL := `create or replace function autorizarCompra(nro_tarjeta char(16), cod_seguridad text, nro_comercio int, monton decimal(7,2)) returns boolean as $$

    declare
        tarjeta_valida record;
        num_operacion int;
        num_rechazo int;
        monto_total float;    
    begin
        num_rechazo := (select coalesce(max(nrorechazo),0) from rechazo);
        select * into tarjeta_valida from tarjeta t where nro_tarjeta = t.nrotarjeta and t.estado = 'vigente'; -- guardo la tarjeta en la variable
    
        if not found tarjeta_valida then
            -- rechazo(nrorechazo int, nro_tarjeta char(16), nro_comercio int, fecha timestamp, monton decimal(7,2), motivo text)
            insert into rechazo values (num_rechazo+1, nro_tarjeta, nro_comercio, CURRENT_TIMESTAMP, monton, '?tarjeta no válida ó no vigente.');
            return false;
        end if;
    
        --if (tarjeta_valida.validahasta < cast(CURRENT_DATE as text)) then
        if (CAST(CONCAT(SUBSTRING(tarjeta_valida.validahasta,0,5),'/',SUBSTRING(tarjeta_valida.validahasta,5,6),'/01') AS DATE) < CURRENT_DATE) then
            insert into rechazo values(num_rechazo+1, nro_tarjeta, nro_comercio, CURRENT_TIMESTAMP, monton,'Tarjeta vencida');
            update tarjeta set estado='anulada' where nrotarjeta=nro_tarjeta;
            return false;
        end if;

        if tarjeta_valida.codseguridad != cod_seguridad then
            insert into rechazo values (num_rechazo+1, nro_tarjeta, nro_comercio, CURRENT_TIMESTAMP, monton,'código de seguridad inválido.');
            return false;
        end if;    
        
        monto_total := (select sum(monto) from compra c where c.nrotarjeta = nro_tarjeta and c.pagado = false);
    
        if tarjeta_valida.limitecompra < monto_total + monton then
            insert into rechazo values (num_rechazo+1, nro_tarjeta, nro_comercio, CURRENT_TIMESTAMP, monton,'supera límite de tarjeta.');
            return false;
        end if;
    
        if tarjeta_valida.estado = 'anulada' then     
            insert into rechazo values(num_rechazo+1, nro_tarjeta, nro_comercio, CURRENT_TIMESTAMP, monton,'plazo de vigencia expirado.');
            return false;
        end if;
        
        if tarjeta_valida.estado = 'suspendida' then     
            insert into rechazo values(num_rechazo+1, nro_tarjeta, nro_comercio, CURRENT_TIMESTAMP, monton,'la tarjeta se encuentra suspendida.');
            return false;
        end if;
    
        -- compra(nrooperacion int, nro_tarjeta string, nro_comercio int, fechaTimeStamp string, monton float64, pagado bool)
        
        num_operacion := (select coalesce(max(nrooperacion),0) from compra);
        insert into compra values (num_operacion+1, nro_tarjeta, nro_comercio, CURRENT_TIMESTAMP, monton, false);
        return true;
    end;
    $$ language plpgsql;`

	db.ExecuteQuery(funcSQL)
	fmt.Println("CREATE FUNCTION - autorizarCompra(nrotarjeta char(16), codseguridad int, nrocomercio int, monton decimal(7,2))")
}

func FuncionAlertaRechazo() {
	funcSQL := `create or replace function alerta_rechazo() returns trigger as $$
    declare 
        nro_alerta int;    
    begin
        nro_alerta := (select coalesce(max(nroalerta),0) from alerta);
        insert into alerta values (nro_alerta+1, new.nrotarjeta, new.fecha, new.nrorechazo, 0, new.motivo);
        return new;
    end;
    $$ language plpgsql;`
	db.ExecuteQuery(funcSQL)
}

func TriggerAlertaRechazo() {
	funcSQL := `create trigger alerta_rechazo_trg
    after insert on rechazo
    for each row
execute procedure alerta_rechazo();`
	db.ExecuteQuery(funcSQL)
}

/*
• generación del resumen el trabajo práctico deberá contener la lógica que reciba
como parámetros el número de cliente, y el periodo del año, y que guarde en las
tablas que corresponda los datos del resumen con la siguiente información: nombre
y apellido, dirección, número de tarjeta, periodo del resumen, fecha de vencimiento,
todas las compras del periodo, y total a pagar.
*/

func TestResumen() {
	// nrocliente - mes - año
	db.ExecuteQuery(`select generarResumen(18, 6, 2022);`)
	db.ExecuteQuery(`select generarResumen(1, 6, 2022);`)
	db.ExecuteQuery(`select generarResumen(15, 4, 2022);`)
}

func FuncionGenerarResumen() {

	funcSQL := `create or replace function generarResumen(num_cliente int, mesP int, anioP int) returns void as $$
declare
    numerotarjeta text;
    nro_resumen int;
    cierre_tarjeta record;
    datoscliente record;
    monto_total decimal;
    cantcompras int;
    datoscomercio record;
    i int;

begin
    select t.nrotarjeta into numerotarjeta from tarjeta t, cliente c where t.nrocliente=c.nrocliente and num_cliente=c.nrocliente;

    select * into cierre_tarjeta from cierre c where c.mes = mesP and c.anio = anioP and terminacion = substring(numerotarjeta,16)::int;

    select nombre,apellido,domicilio into datoscliente from cliente c,tarjeta t where num_cliente=c.nrocliente and t.nrocliente=num_cliente and t.nrotarjeta=numerotarjeta;

    select count(nrooperacion) into cantcompras from compra c where c.nrotarjeta=numerotarjeta;

    for i in 1..cantcompras loop
        select sum(co.monto) into monto_total from compra co;
    end loop;
    
    nro_resumen := (select coalesce(max(nroresumen),0) from cabecera) + 1;

    insert into cabecera values(nro_resumen, datoscliente.nombre,datoscliente.apellido,datoscliente.domicilio,
        numerotarjeta,cierre_tarjeta.fechainicio,cierre_tarjeta.fechacierre,cierre_tarjeta.fechavto, monto_total);

    for i in 1..cantcompras loop
        select c.nombre,co.fecha,co.monto into datoscomercio from comercio c,compra co
            where co.nrooperacion=i and c.nrocomercio=co.nrocomercio and co.nrotarjeta=numerotarjeta;
        insert into detalle values(nro_resumen,i,datoscomercio.fecha,datoscomercio.nombre,datoscomercio.monto);
    end loop;

end;
$$language plpgsql;`
	db.ExecuteQuery(funcSQL)
}

/*
• alertas a clientes el trabajo práctico deberá proveer la lógica que genere alertas por
posibles fraudes. Existe un Call Centre que ante cada alerta generada automáticamente,
realiza un llamado telefónico a le cliente, indicándole la alerta detectada, y
verifica si se trató de un fraude ó no. Se supone que la detección de alertas se ejecuta
automáticamente con cierta frecuencia—e.g. de una vez por minuto. Se pide detectar
y almacenar las siguientes alertas:
– Todo rechazo se debe ingresar automáticamente a la tabla de alertas. No puede
haber ninguna demora para ingresar un rechazo en la tabla de alertas, se debe
ingresar en el mismo instante en que se generó el rechazo.
– Si una tarjeta registra dos compras en un lapso menor de un minuto en comercios
distintos ubicados en el mismo código postal.
– Si una tarjeta registra dos compras en un lapso menor de 5 minutos en comercios
con diferentes códigos postales.
– Si una tarjeta registra dos rechazos por exceso de límite en el mismo día, la tarjeta
tiene que ser suspendida preventivamente, y se debe grabar una alerta asociada a
este cambio de estado.
*/

func FuncionAlertaCompra() {
	funcSQL := `create or replace function alerta_compra() returns trigger as $$
    declare 
        nro_alerta int; 
        compra_anterior record;   
    begin
        nro_alerta := (select coalesce(max(nroalerta),0) from alerta);
        
        select * into compra_anterior from compra c, comercio co 
        where c.nrotarjeta = new.nrotarjeta and c.nrocomercio = co.nrocomercio and c.nrocomercio != new.nrocomercio
        and co.codigopostal = (select codigopostal from comercio com where com.nrocomercio = new.nrocomercio)  
        and c.fecha > CURRENT_TIMESTAMP - interval '1 minute';

        if found then
            insert into alerta values (nro_alerta+1, new.nrotarjeta, new.fecha, null, 1, 'Dos compras en un lapso menor de un minuto');
            return new;
        end if;
        
        -- compra 5 minutos
        select * into compra_anterior from compra c, comercio co 
        where c.nrotarjeta = new.nrotarjeta and c.nrocomercio = co.nrocomercio and c.nrocomercio != new.nrocomercio
        and co.codigopostal != (select codigopostal from comercio com where com.nrocomercio = new.nrocomercio)  
        and c.fecha > CURRENT_TIMESTAMP - interval '5 minute';
        
        if found then 
            insert into alerta values (nro_alerta+1, new.nrotarjeta, new.fecha, null, 5, 'Dos compras en un lapso menor de cinco minutos');
            return new;
        end if;
        return new;
    end;
    $$ language plpgsql;`
	db.ExecuteQuery(funcSQL)
}

func TriggerAlertaCompra() {
	funcSQL := `create trigger alerta_compra_trg
    after insert on compra
    for each row
    execute procedure alerta_compra();`
	db.ExecuteQuery(funcSQL)
}

/** – Si una tarjeta registra dos rechazos por exceso de límite en el mismo día, la tarjeta
tiene que ser suspendida preventivamente, y se debe grabar una alerta asociada a
este cambio de estado. */
func FuncionAlertaLimite() {
	funcSQL := `create or replace function alerta_limite() returns trigger as $$
    declare
        nro_alerta int; 
		cant_rechazos int;
	begin
        nro_alerta := (select coalesce(max(nroalerta),0) from alerta);

		if new.motivo = 'supera límite de tarjeta.' then
			select count(*) into cant_rechazos from rechazo where new.nrotarjeta = rechazo.nrotarjeta and 
            
            --   :: -> PostgreSQL-style typecast, otra forma de castear datos en psql
            
            new.fecha::date = rechazo.fecha::date and rechazo.motivo = 'supera límite de tarjeta.';
			
            if cant_rechazos = 2 then
				update tarjeta set estado = 'suspendida' where new.nrotarjeta = tarjeta.nrotarjeta;
				insert into alerta values(nro_alerta+1, new.nrotarjeta, new.fecha, new.nrorechazo, 32, 'Exceso de limite en el mismo dia');
			end if;
		end if;
		return new;
	end;
    $$ language plpgsql;`
	db.ExecuteQuery(funcSQL)
}

func TriggerAlertaLimite() {
	funcSQL := `create trigger alerta_limite_trg
    after insert on rechazo
    for each row
    execute procedure alerta_limite();`
	db.ExecuteQuery(funcSQL)
}
