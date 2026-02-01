package models

import (
	"time"

)

type Usuario struct {
	ID            uint           `gorm:"primaryKey;column:id_usuario" json:"id"`
	Username      string         `gorm:"size:50;unique;not null;column:username" json:"username"`
	PasswordHash  string         `gorm:"not null;column:password_hash" json:"password,omitempty"` 
	Email         string         `gorm:"size:100;unique;not null;column:email" json:"email"`
	IDRol         uint           `gorm:"column:id_rol" json:"id_rol"`
	UsuarioRol    UsuarioRol     `gorm:"foreignKey:IDRol" json:"rol"`
	FechaCreacion time.Time      `gorm:"column:fecha_creacion;default:CURRENT_TIMESTAMP" json:"fecha_creacion"`
}

func (Usuario) TableName() string { return "usuarios" }

type UsuarioRol struct {
	ID          uint   `gorm:"primaryKey;column:id_rol" json:"id"`
	NombreRol   string `gorm:"size:50;not null;column:nombre_rol" json:"nombre_rol"`
	Descripcion string `gorm:"column:descripcion" json:"descripcion"`
}

func (UsuarioRol) TableName() string { return "usuarios_rol" }

type CargoStaff struct {
	ID          uint   `gorm:"primaryKey;column:id_cargo" json:"id"`
	TituloCargo string `gorm:"size:100;not null;column:titulo_cargo" json:"titulo_cargo"`
	NivelAcceso int    `gorm:"default:1;column:nivel_acceso" json:"nivel_acceso"`
}

func (CargoStaff) TableName() string { return "cargos_staff" }

type StaffAnimal struct {
	ID         uint       `gorm:"primaryKey;column:id_staff" json:"id"`
	Nombre     string     `gorm:"size:100;not null;column:nombre" json:"nombre"`
	Especie    string     `gorm:"size:50;column:especie" json:"especie"`
	AvatarURL  string     `gorm:"column:avatar_url" json:"avatar_url"`
	IDCargo    uint       `gorm:"column:id_cargo" json:"id_cargo"`
	CargoStaff CargoStaff `gorm:"foreignKey:IDCargo" json:"cargo"`
}

func (StaffAnimal) TableName() string { return "staff_animal" }

type FamiliaBotanica struct {
	ID            uint   `gorm:"primaryKey;column:id_familia" json:"id"`
	NombreFamilia string `gorm:"size:100;not null;column:nombre_familia" json:"nombre_familia"`
	Descripcion   string `gorm:"column:descripcion" json:"descripcion"`
}

func (FamiliaBotanica) TableName() string { return "familias_botanicas" }

type EstadoSalud struct {
	ID           uint   `gorm:"primaryKey;column:id_estado" json:"id"`
	NombreEstado string `gorm:"size:50;not null;column:nombre_estado" json:"nombre_estado"`
	ColorHex     string `gorm:"size:7;column:color_hex" json:"color_hex"`
}

func (EstadoSalud) TableName() string { return "estados_salud" }

type Especie struct {
    IDEspecie        uint            `gorm:"primaryKey;column:id_especie" json:"id"`
    NombreComun      string          `gorm:"size:100;not null;column:nombre_comun" json:"nombre_comun"`
    NombreCientifico string          `gorm:"size:100;column:nombre_cientifico" json:"nombre_cientifico"`
    IDFamilia        uint            `gorm:"column:id_familia" json:"id_familia"`
    FamiliaBotanica  FamiliaBotanica `gorm:"foreignKey:IDFamilia;references:ID" json:"familia"`
    IDEstado         uint            `gorm:"column:id_estado" json:"id_estado"`
    EstadoSalud      EstadoSalud     `gorm:"foreignKey:IDEstado;references:ID" json:"estado"`
    FechaRegistro    time.Time       `gorm:"column:fecha_registro;default:CURRENT_TIMESTAMP" json:"fecha_registro"`
    ImagenURL        string          `gorm:"column:imagen_url" json:"imagen_url"`
}

func (Especie) TableName() string { return "especies" }

type RegistroSalud struct {
	ID                uint      `gorm:"primaryKey;column:id_registro" json:"id"`
	IDEspecie         uint      `gorm:"column:id_especie" json:"id_especie"`
	Especie           Especie   `gorm:"foreignKey:IDEspecie" json:"especie"`
	HumedadPorcentaje float64   `gorm:"type:numeric(5,2);column:humedad_porcentaje" json:"humedad"`
	TempCelsius       float64   `gorm:"type:numeric(5,2);column:temperatura_celsius" json:"temperatura"`
	FechaMonitoreo    time.Time `gorm:"column:fecha_monitoreo;default:CURRENT_TIMESTAMP" json:"fecha_monitoreo"`
}

func (RegistroSalud) TableName() string { return "registros_salud" }