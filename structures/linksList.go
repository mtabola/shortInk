package structures

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type LinksList struct {
	Links []Link
}

func (l *LinksList) FillFromDatabase(db *sql.DB) error {
	res, err := db.Query("SELECT * FROM link")

	if err != nil {
		return err
	}

	for res.Next() {
		var lk Link
		err = res.Scan(&lk.LinkId, &lk.FullLink, &lk.ShortLink)
		if err != nil {
			return err
		}
		l.Links = append(l.Links, lk)
	}
	return nil
}

func (l LinksList) FindLinkByShortHash(sl string) *Link {
	for _, link := range l.Links {
		if link.ShortLink == sl {
			return &link
		}
	}
	return nil
}

func (l LinksList) GetAllLinks() []Link {
	return l.Links
}

func (l *LinksList) AddLink(link Link, db *sql.DB) error {
	for _, lk := range l.Links {
		if link.LinkId == lk.LinkId {
			return errors.New("NotUniqueId")
		} else if link.ShortLink == lk.ShortLink {
			return errors.New("NotUniqueShortLink")
		}
	}
	l.Links = append(l.Links, link)

	_, err := db.Exec("INSERT INTO link(FullLink, ShortLink) VALUES (?, ?)", link.FullLink, link.ShortLink)

	if err != nil {
		return errors.New("Insert in database fault")
	}

	return nil
}

func (l *LinksList) DeleteLink(link Link, db *sql.DB) error {
	for i, lk := range l.Links {
		if lk.LinkId == link.LinkId {
			_, err := db.Exec("DELETE FROM link WHERE ShortLink = ?", link.ShortLink)
			l.Links = append(l.Links[:i], l.Links[i+1:]...)
			if err != nil {
				log.Print(fmt.Sprintf(err.Error()))
			}

			return nil
		}
	}
	return errors.New("This link doesn't exist")
}
