-- MySQL dump 10.13  Distrib 5.7.9, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: starwars
-- ------------------------------------------------------
-- Server version	5.7.12-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
--
-- Table structure for table `character_friend`
--

DROP TABLE IF EXISTS `character_friend`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `character_friend` (
  `charid` varchar(10) NOT NULL,
  `friendid` varchar(10) NOT NULL,
  PRIMARY KEY (`charid`,`friendid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `character_friend`
--

LOCK TABLES `character_friend` WRITE;
/*!40000 ALTER TABLE `character_friend` DISABLE KEYS */;
INSERT INTO `character_friend` VALUES ('1000','1002'),('1000','1003'),('1000','2000'),('1000','2001');
/*!40000 ALTER TABLE `character_friend` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `character_type`
--

DROP TABLE IF EXISTS `character_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `character_type` (
  `id` int(11) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `character_type`
--

LOCK TABLES `character_type` WRITE;
/*!40000 ALTER TABLE `character_type` DISABLE KEYS */;
INSERT INTO `character_type` VALUES (1000,'Human'),(2000,'Droid');
/*!40000 ALTER TABLE `character_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `starwarschar`
--

DROP TABLE IF EXISTS `starwarschar`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `starwarschar` (
  `id` varchar(10) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `home_planet` varchar(45) DEFAULT NULL,
  `primary_function` varchar(45) DEFAULT NULL,
  `type_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `char_type_idx` (`type_id`),
  CONSTRAINT `char_type` FOREIGN KEY (`type_id`) REFERENCES `character_type` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `starwarschar`
--

LOCK TABLES `starwarschar` WRITE;
/*!40000 ALTER TABLE `starwarschar` DISABLE KEYS */;
INSERT INTO `starwarschar` VALUES ('1000','Luke Skywalker','Tatooine',NULL,1000),('1001','Darth Vader','Tatooine',NULL,1000),('1002','Han Solo',NULL,NULL,1000),('1003','Leia Organa','Alderaan',NULL,1000),('1004','Wilhuff Tarkin',NULL,NULL,1000),('2000','C-3PO',NULL,'Protocol',2000),('2001','R2-D2',NULL,'Astromech',2000);
/*!40000 ALTER TABLE `starwarschar` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-11-17 15:09:10
