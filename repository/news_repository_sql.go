package repository

const (
	GetNewsByUrl = `
		SELECT 
			IF(COUNT(1), 'true', 'false') 
		FROM 
			news_article 
		WHERE 
			url = ?
	`

	CreateNews = `
		INSERT INTO news_article (
			title,
			url,
			content,
			summary,
			article_ts,
			published_date,
			inserted,
			updated
		) VALUES (?,?,?,?,?,?,?,?)
	`

	FindNewsByID = `
		SELECT 
			id, title, url, content 
		FROM 
			news_article 
		WHERE 
			id = ?
	`
)
