package content

const (
	getByTopicIDsStmt = `SELECT id, topic_id, content, additional_resources, video_urls, image_urls, created_at, updated_at 
						FROM topic_contents 
						WHERE topic_id = ANY ($1)`
)
