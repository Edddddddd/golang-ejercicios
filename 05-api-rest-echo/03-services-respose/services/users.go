package services

import (
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type Usuarios struct {
	ID int32 `json:"id"`
	IdGorilas string `json:"idGorilas"`
	NombreCompleto string `json:"nombreCompleto"`
	CorreoElectronico string `json:"correoElectronico"`
	Documento string `json:"documento"`
	Abono int64 `json:"abono"`
	Estado string `json:"estado"`
	TurnosDisponibles int32 `json:"turnosDisponibles"`
	TurnosDisponiblesLeyenda string `json:"turnosDisponiblesLeyenda"`
	TurnosDisponiblesLeyendaHtml string `json:"turnosDisponiblesLeyendaHtml"`
	AbonosVencimientoHtml string `json:"abonosVencimientoHtml"`
	FechaProxAptoFisico time.Time `json:"fechaProxAptoFisico"`
	FechaProxAptoFisicoString time.Time `json:"fechaProxAptoFisicoString"`
	DiasRestantes float32 `json:"diasRestantes"`
	TieneDeuda bool `json:"tieneDeuda"`
	Saldo int32 `json:"saldo"`
	Domicilios string `json:"domicilio"`
	Barrio string `json:"barrio"`
}


// Consulta informacion de los estudiantes
func ConsultarUsuarioGorilas() (usuarios []Usuarios, err error) {
	q := `SELECT 
				id ,
				id_gorilas ,
				nombre_completo ,
				correo_electronico ,
				documento ,
				abono ,
				estado ,
				turnos_disponibles ,
				turnos_disponibles_leyenda ,
				turnos_disponibles_leyenda_html ,
				abonos_vencimiento_html ,
				fecha_prox_apto_fisico ,
				fecha_prox_apto_fisico_string ,
				dias_restantes ,
				tiene_deuda ,
				saldo ,
				domicilio ,
				barrio 
 			FROM
			public.usuarios`


	id_Null := sql.NullInt64{}
	id_gorilas_Null := sql.NullString{}
	nombre_completo_Null := sql.NullString{}
	correo_electronico_Null := sql.NullString{}
	documento_Null := sql.NullString{}
	abono_Null := sql.NullInt64{}
	estado_Null := sql.NullString{}
	turnos_disponibles_Null := sql.NullInt64{}
	turnos_disponibles_leyenda_Null := sql.NullString{}
	turnos_disponibles_leyenda_html_Null := sql.NullString{}
	abonos_vencimiento_html_Null := sql.NullString{}
	fecha_prox_apto_fisico_Null := pq.NullTime{}
	fecha_prox_apto_fisico_string_Null := pq.NullTime{}
	dias_restantes_Null := sql.NullFloat64{}
	tiene_deuda_Null := sql.NullBool{}
	saldo_Null := sql.NullInt64{}
	domicilio_Null := sql.NullString{}
	barrio_Null := sql.NullString{}

	db := a.GetConnection()
	defer db.Close()
	rows, err := db.Query(q)
	if err != nil {
		return nil, nil
	}
	defer rows.Close()

	for rows.Next() {
		e := Usuarios{}
		err = rows.Scan(
			&id_Null ,
			&id_gorilas_Null ,
			&nombre_completo_Null ,
			&correo_electronico_Null ,
			&documento_Null ,
			&abono_Null ,
			&estado_Null ,
			&turnos_disponibles_Null ,
			&turnos_disponibles_leyenda_Null ,
			&turnos_disponibles_leyenda_html_Null ,
			&abonos_vencimiento_html_Null ,
			&fecha_prox_apto_fisico_Null ,
			&fecha_prox_apto_fisico_string_Null ,
			&dias_restantes_Null ,
			&tiene_deuda_Null ,
			&saldo_Null ,
			&domicilio_Null ,
			&barrio_Null,
		)
		if err != nil {
			return nil, nil
		}

		e.ID = int32(id_Null.Int64)
		e.IdGorilas = id_gorilas_Null.String
		e.NombreCompleto = nombre_completo_Null.String
		e.CorreoElectronico = correo_electronico_Null.String
		e.Documento = documento_Null.String
		e.Abono = abono_Null.Int64
		e.Estado= estado_Null.String
		e.TurnosDisponibles = int32(turnos_disponibles_Null.Int64)
		e.TurnosDisponiblesLeyenda = turnos_disponibles_leyenda_Null.String
		e.TurnosDisponiblesLeyendaHtml = turnos_disponibles_leyenda_html_Null.String
		e.AbonosVencimientoHtml = abonos_vencimiento_html_Null.String
		e.FechaProxAptoFisico = fecha_prox_apto_fisico_Null.Time
		e.FechaProxAptoFisicoString = fecha_prox_apto_fisico_string_Null.Time
		e.DiasRestantes = float32(dias_restantes_Null.Float64)
		e.TieneDeuda = tiene_deuda_Null.Bool
		e.Saldo = int32(saldo_Null.Int64)
		e.Domicilios= domicilio_Null.String
		e.Barrio = barrio_Null.String

		usuarios = append(usuarios, e)
	}
	return usuarios, nil

}