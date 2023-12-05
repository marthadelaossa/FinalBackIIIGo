CREATE DATABASE IF NOT EXISTS `clinicaodontologica` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE TABLE IF NOT EXISTS `odontologo` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `last_name` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `medical_ID` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `paciente` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `last_name` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `DNI` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `address` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `creation_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE  IF NOT EXISTS `appointments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(45) COLLATE utf8mb4_general_ci NOT NULL,
  `id_odontologo` int NOT NULL,
  `id_paciente` int NOT NULL,
  `date_time` DATETIME,
  PRIMARY KEY (`id`),
  KEY `id_odontologo` (`id_odontologo`),
  KEY `id_paciente` (`id_paciente`),
  CONSTRAINT `id_odontologo` FOREIGN KEY (`id_odontologo`) REFERENCES `odontologo` (`id`),
  CONSTRAINT `id_paciente` FOREIGN KEY (`id_paciente`) REFERENCES `paciente` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;