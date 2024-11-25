BACKUP_FILE="/tmp/dump.sql"
DEST_PARENT_DIR="/tmp/dump"
LOG_FILE="/tmp/dump/backup.log"
MAX_BACKUPS=100

TIMESTAMP=$(date +"%Y%m%d%H%M%S")
cp ${BACKUP_FILE} ${DEST_PARENT_DIR}/dump_${TIMESTAMP}.sql
mysqldump -h 1.2.3.4 -P 3306 --databases biz_db > ${BACKUP_FILE}

BACKUP_COUNT=$(ls -1t ${DEST_PARENT_DIR}/dump_* | wc -l)

if [ "$BACKUP_COUNT" -gt "$MAX_BACKUPS" ]; then
    echo "Backup count exceeded the threshold ($MAX_BACKUPS). Deleting the oldest backups." >> "$LOG_FILE"

    ls -1t ${DEST_PARENT_DIR}/dump_* | tail -n +$(($MAX_BACKUPS + 1)) | xargs rm -rf

    echo "Oldest backups deleted." >> "$LOG_FILE"
fi
