CREATE TABLE blog_to_tag(
    blog_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    FOREIGN KEY (blog_id) REFERENCES blog(id),
    FOREIGN KEY (tag_id) REFERENCES tag(id),
    PRIMARY KEY (blog_id, tag_id)
);