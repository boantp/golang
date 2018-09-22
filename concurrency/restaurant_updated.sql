/*
SQLyog Ultimate v10.42 
MySQL - 5.5.5-10.1.28-MariaDB : Database - restaurant
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`restaurant` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `restaurant`;

/*Table structure for table `cuisines` */

DROP TABLE IF EXISTS `cuisines`;

CREATE TABLE `cuisines` (
  `cuisines_id` int(11) NOT NULL AUTO_INCREMENT,
  `cuisines_name` varchar(100) NOT NULL,
  PRIMARY KEY (`cuisines_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;

/*Data for the table `cuisines` */

insert  into `cuisines`(`cuisines_id`,`cuisines_name`) values (1,'chinese'),(2,'indian'),(3,'indonesian'),(4,'italian'),(5,'japanese'),(6,'korean'),(7,'pizza'),(8,'seafood'),(9,'sushi'),(10,'french');

/*Table structure for table `customer` */

DROP TABLE IF EXISTS `customer`;

CREATE TABLE `customer` (
  `customer_id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_name` varchar(100) NOT NULL,
  `customer_phone` varchar(30) NOT NULL,
  `customer_email` varchar(100) NOT NULL,
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `customer` */

insert  into `customer`(`customer_id`,`customer_name`,`customer_phone`,`customer_email`) values (1,'Boan','085210549044','boantuapasaribu@gmail.com');

/*Table structure for table `operational` */

DROP TABLE IF EXISTS `operational`;

CREATE TABLE `operational` (
  `operational_id` int(11) NOT NULL AUTO_INCREMENT,
  `operational_restaurant_id` int(11) NOT NULL,
  `operational_day` varchar(20) NOT NULL,
  `operational_open_hour` varchar(50) NOT NULL,
  `operational_closed_hour` varchar(50) NOT NULL,
  PRIMARY KEY (`operational_id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=latin1;

/*Data for the table `operational` */

insert  into `operational`(`operational_id`,`operational_restaurant_id`,`operational_day`,`operational_open_hour`,`operational_closed_hour`) values (1,1,'Saturday','10:30AM','10:30PM'),(2,1,'Sunday','10:30AM','10:30PM'),(3,1,'Monday','10:30AM','10:30PM'),(4,1,'Tuesday','10:30AM','10:30PM'),(5,1,'Wednesday','10:30AM','10:30PM'),(6,1,'Thursday','10:30AM','10:30PM'),(7,1,'Friday','10:30AM','10:30PM'),(8,2,'Saturday','10:30AM','10:30PM'),(9,2,'Sunday','10:30AM','10:30PM'),(10,2,'Monday','10:30AM','10:30PM'),(11,2,'Tuesday','10:30AM','10:30PM'),(12,2,'Wednesday','10:30AM','10:30PM'),(13,2,'Thursday','10:30AM','10:30PM'),(14,2,'Friday','10:30AM','10:30PM'),(15,3,'Saturday','10:30AM','10:30PM'),(16,3,'Sunday','10:30AM','10:30PM'),(17,3,'Monday','10:30AM','10:30PM'),(18,3,'Tuesday','10:30AM','10:30PM'),(19,3,'Wednesday','10:30AM','10:30PM'),(20,3,'Thursday','10:30AM','10:30PM'),(21,3,'Friday','10:30AM','10:30PM');

/*Table structure for table `reservation` */

DROP TABLE IF EXISTS `reservation`;

CREATE TABLE `reservation` (
  `reservation_id` int(11) NOT NULL AUTO_INCREMENT,
  `reservation_restaurant_id` int(11) NOT NULL,
  `reservation_code` varchar(50) NOT NULL,
  `reservation_total_guest` varchar(50) NOT NULL,
  `reservation_datetime` datetime NOT NULL,
  `reservation_customer_id` int(11) NOT NULL,
  `reservation_customer_name` varchar(50) NOT NULL,
  `reservation_customer_phone` varchar(30) NOT NULL,
  PRIMARY KEY (`reservation_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `reservation` */

insert  into `reservation`(`reservation_id`,`reservation_restaurant_id`,`reservation_code`,`reservation_total_guest`,`reservation_datetime`,`reservation_customer_id`,`reservation_customer_name`,`reservation_customer_phone`) values (1,1,'1234567','2','2018-09-22 10:15:28',1,'Boan','085210549044');

/*Table structure for table `restaurant` */

DROP TABLE IF EXISTS `restaurant`;

CREATE TABLE `restaurant` (
  `restaurant_id` int(11) NOT NULL AUTO_INCREMENT,
  `restaurant_name` varchar(100) NOT NULL,
  `restaurant_description` text NOT NULL,
  `restaurant_address` text NOT NULL,
  `restaurant_phone` varchar(30) NOT NULL,
  `restaurant_location` varchar(100) NOT NULL,
  `restaurant_cuisines_id` int(11) NOT NULL,
  `restaurant_latitude` double NOT NULL,
  `restaurant_longitude` double NOT NULL,
  `restaurant_image` varchar(255) NOT NULL,
  PRIMARY KEY (`restaurant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

/*Data for the table `restaurant` */

insert  into `restaurant`(`restaurant_id`,`restaurant_name`,`restaurant_description`,`restaurant_address`,`restaurant_phone`,`restaurant_location`,`restaurant_cuisines_id`,`restaurant_latitude`,`restaurant_longitude`,`restaurant_image`) values (1,'Emilie French Restaurant','EMILIE OFFERS AN ARTFUL BLEND OF MODERN YET CLASSICALLY INSPIRED FRENCH CUISINE. SIMPLE, YET HIGH ATTENTION TO DETAIL, EACH DISH BENEFITS FROM THE BEST OF \r\nSEASONAL PRODUCE. ','Jalan Senopati 39, Jakarta Capital Region, RT.6/RW.3, Senayan, Kby. Baru, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12190','(021)5213626','Jakarta',10,-6.231031,106.809259,'french_jakarta.jpg'),(2,'Mikawa (Japanese Restaurant)','Mikawa Japanese Sake Bar & Restaurant offers what Japan cuisine has to brag to their food about with its mouth-watering dishes and delectable foods that could satisfy your cravings with variety of selection of Sake, Sashimi, Sushi, Tempura, Yakitori and many more. This is a place to rest your mind and eat to your satisfaction. .','Jl. Sukaraja II No.32, Sukaraja, Cicendo, Kota Bandung, Jawa Barat 40175','(022)6073983','Bandung',5,-6.892334,107.574673,'japan_bandung.jpg'),(3,'Saung Kuring Sundanese Restaurant','Sundanese traditional food from indonesia','Jalan KH Sholeh Iskandar No.9, Kedung Badak, Tanah Sereal, Kedung Badak, Tanah Sereal, Kota Bogor, Jawa Barat 16164','(0251)8331885','Bogor',3,-6.561851,106.799155,'indonesia_bogor.jpg');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
