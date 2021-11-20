-- MySQL dump 10.13  Distrib 5.7.35, for Linux (x86_64)
--
-- Host: localhost    Database: indian_geography_quiz
-- ------------------------------------------------------
-- Server version	5.7.35-0ubuntu0.18.04.2

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
-- Table structure for table `entries`
--

DROP TABLE IF EXISTS `entries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `entries` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `quiz_id` int(11) NOT NULL,
  `quizsubcategory_id` int(11) NOT NULL,
  `entry` varchar(255) NOT NULL,
  `image_link` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=29 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `entries`
--

LOCK TABLES `entries` WRITE;
/*!40000 ALTER TABLE `entries` DISABLE KEYS */;
INSERT INTO `entries` VALUES (1,1,1,'Andhra Pradesh','states/andhra.png'),(2,1,1,'Arunachal Pradesh','states/arunachal.png'),(3,1,1,'Assam','states/assam.png'),(4,1,1,'Bihar','states/bihar.png'),(5,1,1,'Chattisgarh','states/chattisgarh.png'),(6,1,1,'Goa','states/goa.png'),(7,1,1,'Gujrat','states/gujrat.png'),(8,1,1,'Haryana','states/haryana.png'),(9,1,1,'Himachal Pradesh','states/hp.png'),(10,1,1,'Jharkhand','states/jharkhand.png'),(11,1,1,'Karnataka','states/karnataka.png'),(12,1,1,'Kerala','states/kerala.png'),(13,1,1,'Maharashtra','states/maharashtra.png'),(14,1,1,'Manipur','states/manipur.png'),(15,1,1,'Meghalaya','states/mehalaya.png'),(16,1,1,'Mizoram','states/mizoram.png'),(17,1,1,'Madhya Pradesh','states/mp.png'),(18,1,1,'Nagaland','states/nagaland.png'),(19,1,1,'Odisha','states/odhisha.png'),(20,1,1,'Punjab','states/punjab.png'),(21,1,1,'Rajasthan','states/raj.png'),(22,1,1,'Sikkim','states/sikkim.png'),(23,1,1,'Telanana','states/telangana.png'),(24,1,1,'Tamil Nadu','states/tn.png'),(25,1,1,'Tripura','states/tripura.png'),(26,1,1,'Uttarakhand','states/uk.png'),(27,1,1,'Uttar Pradesh','states/up.png'),(28,1,1,'West Bengal','states/web.png');
/*!40000 ALTER TABLE `entries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `leaderboards`
--

DROP TABLE IF EXISTS `leaderboards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `leaderboards` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `score` smallint(6) NOT NULL,
  `quiz_id` int(11) NOT NULL,
  `score_date` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `leaderboards`
--

LOCK TABLES `leaderboards` WRITE;
/*!40000 ALTER TABLE `leaderboards` DISABLE KEYS */;
INSERT INTO `leaderboards` VALUES (1,1,4,1,'2021-11-15'),(2,2,3,1,'2021-11-15'),(3,3,2,1,'2021-11-16'),(4,2,1,1,'2021-11-16'),(5,1,6,1,'2021-11-16'),(6,1,10,1,'2021-11-17'),(7,1,19,1,'2021-11-18');
/*!40000 ALTER TABLE `leaderboards` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `quizes`
--

DROP TABLE IF EXISTS `quizes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `quizes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `quiz_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `quizes`
--

LOCK TABLES `quizes` WRITE;
/*!40000 ALTER TABLE `quizes` DISABLE KEYS */;
INSERT INTO `quizes` VALUES (1,'States');
/*!40000 ALTER TABLE `quizes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `quizsubcategories`
--

DROP TABLE IF EXISTS `quizsubcategories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `quizsubcategories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `quiz_id` int(11) NOT NULL,
  `subcategory_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `quizsubcategories`
--

LOCK TABLES `quizsubcategories` WRITE;
/*!40000 ALTER TABLE `quizsubcategories` DISABLE KEYS */;
INSERT INTO `quizsubcategories` VALUES (1,1,'All States');
/*!40000 ALTER TABLE `quizsubcategories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `session_key` varchar(255) DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `joined_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `email` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'$2a$14$uDt74gqu6QYQtodbFi897.eObOodwZEG0n3qjvr4dUD5LwSs09n7u','asd1','2021-11-15 04:48:44','chilly5476@gmail.com'),(2,'sdfsdqwerr','meenakshi','2021-11-15 04:48:44','capricioussharan@gmail.com'),(3,'adsasd','Vishal Saharan','2021-11-16 08:58:32','vishal.sahara33@gmail.com'),(7,'RN6SeP7SLfUHywfKf5gJGJhLgC3DdPaSJGW','x1','2021-11-19 09:16:53','x12@gmail.com'),(8,'$2a$14$8aU4/EsG0uleSomxhMHwa.mkk7C/DTZUJ5nsjkf39pEm.ReP8ACeG','vishal.saharan33@gmail.com','2021-11-19 11:41:54','vishal.saharan223@gmail.com');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-11-21  1:41:09
