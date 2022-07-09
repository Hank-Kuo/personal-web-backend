CREATE TABLE blog(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    link VARCHAR(255) NOT NULL,
    visitor INTEGER NOT NULL DEFAULT 0,
    img_link VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
    blog (title, link, img_link)
VALUES
    (
        "自然語言處理 — 使用 N-gram 實現輸入文字預測",
        "// https://github.com/Hank-Kuo/blog/blob/master/NLP.md",
        "https://i.imgur.com/5ezNa4m.jpg"
    )