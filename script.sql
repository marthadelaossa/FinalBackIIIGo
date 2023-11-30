CREATE DATABASE IF NOT EXISTS `clinicaOdontologica` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE TABLE IF NOT EXISTS `odontologo` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `last_name` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `matricula` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `paciente` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `last_name` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `DNI` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `domicilio` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `fecha_alta` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE  IF NOT EXISTS `turnos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(45) COLLATE utf8mb4_general_ci NOT NULL,
  `id_odontologo` int NOT NULL,
  `id_paciente` int NOT NULL,
  `fecha_hora` DATETIME,
  `descripcion` VARCHAR(255),
  PRIMARY KEY (`id`),
  KEY `id_odontologo` (`id_odontologo`),
  KEY `id_paciente` (`id_paciente`),
  CONSTRAINT `id_odontologo` FOREIGN KEY (`id_odontologo`) REFERENCES `odontologo` (`id`),
  CONSTRAINT `id_paciente` FOREIGN KEY (`id_paciente`) REFERENCES `paciente` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8