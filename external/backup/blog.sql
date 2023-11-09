CREATE TABLE `news` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `title` varchar(120) NOT NULL,
  `content` text NOT NULL,
  `status_content` ENUM('drafted', 'published', 'deleted') NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `topics` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(120) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `ref_news_topics` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `news_id` int(10) NOT NULL,
  `topics_id` int(10) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (news_id) REFERENCES news (id) ON DELETE CASCADE,
  FOREIGN KEY (topics_id) REFERENCES topics (id) ON DELETE CASCADE
);

INSERT INTO `news` (id,title,content,status_content) VALUES
(1,'HTML and CSS code snippets that you can use for your own projects','MJT provides a comprehensive guide on creating various types of three-column layouts using CSS Grid and Flexbox.','deleted'),
(2,'Introduction to IEEE','IEEE Author Center Journals provides article templates to help you format your article and prepare a draft for peer review. The templates help with the placement of specific elements, such as the author list, and provide guidance on stylistic elements such as abbreviations and acronyms','published'),
(3,'Introduction with W3Schools','W3Schools has a tutorial on creating a responsive three-column layout using CSS. The tutorial includes HTML and CSS code snippets that you can use for your own projects.','published'),
(4,'Introduction with Scribbr','Scribbr provides guidelines for formatting IEEE papers, including formatting the text as two columns, in Times New Roman, 10 pt','drafted');

INSERT INTO `topics` (id,category_name) VALUES
(1,'olahraga'),
(2,'berita'),
(3,'paper'),
(4,'book');

INSERT INTO `ref_news_topics` (id,news_id,topics_id) VALUES
(1,1,2),
(2,2,2),
(3,3,2),
(4,4,1);