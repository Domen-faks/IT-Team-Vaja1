package MariaDB

//Začetna verzija baze
var createDB = `
CREATE TABLE IF NOT EXISTS task(
    task_id INT NOT NULL AUTO_INCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    date_added TEXT NOT NULL,
	working_date_estimate TEXT NOT NULL,
    PRIMARY KEY (task_id)
)
CHARACTER SET 'utf8mb4',
COLLATE 'utf8mb4_unicode_ci'
`
