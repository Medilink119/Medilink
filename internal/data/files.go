package data

import (
	"database/sql"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type File struct {
	FileId    int
	UserId    int
	CreatedAt time.Time
	Name      string
	Extension string
	Category  string
	PrevName  string
	Type      string
}

type FileModel struct {
	DB *sql.DB
}

const PRESCRIPTION = "prescription"
const SCAN = "scan"

// ================== FileModel Methods =====================

func (fm *FileModel) Insert(file *File) error {
	query := `INSERT INTO files (user_id, name, ext, cat, prev_name, type) VALUES ($1, $2, $3, $4, $5, $6) RETURNING file_id;`
	args := []interface{}{file.UserId, file.Name, file.Extension, file.Category, file.PrevName, file.Type}
	err := fm.DB.QueryRow(query, args...).Scan(&file.FileId)
	if err != nil {
		return err
	}
	return nil
}

func (fm *FileModel) GetFileById(id int) (*File, error) {
	file := &File{}
	query := `SELECT * FROM files WHERE file_id = ($1)`
	err := fm.DB.QueryRow(query, id).Scan(
		&file.FileId,
		&file.UserId,
		&file.CreatedAt,
		&file.Name,
		&file.Extension,
		&file.Category,
		&file.PrevName,
		&file.Type,
	)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (fm *FileModel) GetCategoriesByUser(userId int, ftype string) ([]string, error) {
	var categories []string
	query := `SELECT DISTINCT cat FROM files WHERE user_id = ($1) and type = ($2)`
	rows, err := fm.DB.Query(query, userId, ftype)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cat string
		rows.Scan(&cat)
		categories = append(categories, cat)
	}

	return categories, nil
}

func (fm *FileModel) GetUserFilesByCategory(userId int, cat, ftype string) ([]File, error) {
	var files []File
	query := `SELECT * FROM files WHERE cat = ($1) and type = ($2) and user_id = ($3)`
	rows, err := fm.DB.Query(query, cat, ftype, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var f File
		rows.Scan(
			&f.FileId,
			&f.UserId,
			&f.CreatedAt,
			&f.Name,
			&f.Extension,
			&f.Category,
			&f.PrevName,
			&f.Type,
		)
		files = append(files, f)
	}

	return files, nil
}

func (fm *FileModel) GetUserFiles(userId int, ftype string) ([]File, error) {
	var files []File
	query := `SELECT * FROM files WHERE user_id = ($1) and type = ($2)`
	rows, err := fm.DB.Query(query, userId, ftype)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var f File
		rows.Scan(
			&f.FileId,
			&f.UserId,
			&f.CreatedAt,
			&f.Name,
			&f.Extension,
			&f.Category,
			&f.PrevName,
			&f.Type,
		)
		files = append(files, f)
	}

	return files, nil
}

// ================== File Methods =========================

// create a new instance of File Struct
// ftype: 'scan/report' or 'prescription'
// category: 'general', 'orthopeodic' etc....
func NewFile(file multipart.File, header *multipart.FileHeader, category string, userId int, ftype string) *File {
	return &File{
		UserId:    userId,
		Name:      uuid.NewString(),
		Extension: filepath.Ext(header.Filename),
		Category:  category,
		PrevName:  header.Filename,
		Type:      ftype,
	}
}
