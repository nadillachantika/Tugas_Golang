/*
SQLyog Community v13.1.7 (64 bit)
MySQL - 10.4.11-MariaDB : Database - perkuliahan
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`perkuliahan` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `perkuliahan`;

/*Table structure for table `alamat_details` */

DROP TABLE IF EXISTS `alamat_details`;

CREATE TABLE `alamat_details` (
  `KodeAlamat` varchar(10) NOT NULL,
  `Jalan` varchar(50) DEFAULT NULL,
  `Kelurahan` varchar(50) DEFAULT NULL,
  `Kecamatan` varchar(50) DEFAULT NULL,
  `KotaKabupaten` varchar(50) DEFAULT NULL,
  `Provinsi` varchar(50) DEFAULT NULL,
  `MahasiswaID` int(11) DEFAULT NULL,
  KEY `kode_alamat` (`KodeAlamat`),
  KEY `id_mahasiswa` (`MahasiswaID`),
  CONSTRAINT `alamat_details_ibfk_1` FOREIGN KEY (`MahasiswaID`) REFERENCES `mahasiswa` (`MahasiswaID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `alamat_details` */

insert  into `alamat_details`(`KodeAlamat`,`Jalan`,`Kelurahan`,`Kecamatan`,`KotaKabupaten`,`Provinsi`,`MahasiswaID`) values 
('BKT','Jl. Panorama Baru','Puhun Pintu Kabun','Mandiangin Koto Selayan','Bukittinggi','Sumater Barat',4),
('BKT','Jl. Kampung Pulasan','Puhun Pintu Kabun','Mandiangin Koto Selayan','Bukittinggi','Sumatera Barat',6),
('PDG','Jl. Koto Panjang','Limau Manis','Pauh','Padang','Sumatera Barat',5);

/*Table structure for table `mahasiswa` */

DROP TABLE IF EXISTS `mahasiswa`;

CREATE TABLE `mahasiswa` (
  `MahasiswaID` int(11) NOT NULL AUTO_INCREMENT,
  `NoBp` varchar(20) DEFAULT NULL,
  `Nama` varchar(50) DEFAULT NULL,
  `Jurusan` varchar(30) DEFAULT NULL,
  `Prodi` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`MahasiswaID`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

/*Data for the table `mahasiswa` */

insert  into `mahasiswa`(`MahasiswaID`,`NoBp`,`Nama`,`Jurusan`,`Prodi`) values 
(4,'1811082013','Nadilla C. Putri','TI','TRPL'),
(5,'1811082014','Fikri','TI','TRPL'),
(6,'1811082090','Anjely',NULL,NULL),
(7,'1811082102','Lailatul Fadhila','TI','MI');

/*Table structure for table `mata_kuliah` */

DROP TABLE IF EXISTS `mata_kuliah`;

CREATE TABLE `mata_kuliah` (
  `KodeMatKul` varchar(10) NOT NULL,
  `NamaMatKul` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`KodeMatKul`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `mata_kuliah` */

insert  into `mata_kuliah`(`KodeMatKul`,`NamaMatKul`) values 
('PMOB','P. Pemprograman Mobile'),
('PWEB','P. Pemprograman WEB'),
('TKAL','Kalkulus'),
('TMPL','');

/*Table structure for table `tabel_nilai` */

DROP TABLE IF EXISTS `tabel_nilai`;

CREATE TABLE `tabel_nilai` (
  `MahasiswaID` int(11) DEFAULT NULL,
  `KodeMatKul` varchar(10) DEFAULT NULL,
  `Nilai` float DEFAULT NULL,
  `Semester` varchar(10) DEFAULT NULL,
  KEY `kode_matakuliah` (`KodeMatKul`),
  KEY `MahasiswaID` (`MahasiswaID`,`KodeMatKul`),
  CONSTRAINT `tabel_nilai_ibfk_2` FOREIGN KEY (`KodeMatKul`) REFERENCES `mata_kuliah` (`KodeMatKul`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `tabel_nilai_ibfk_3` FOREIGN KEY (`MahasiswaID`) REFERENCES `mahasiswa` (`MahasiswaID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `tabel_nilai` */

insert  into `tabel_nilai`(`MahasiswaID`,`KodeMatKul`,`Nilai`,`Semester`) values 
(4,'PWEB',3.75,'4'),
(4,'PMOB',3.5,'4'),
(5,'TKAL',3.5,'2');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
