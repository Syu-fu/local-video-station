INSERT INTO video (id, title, title_reading, url, thumbnail_url)
VALUES
    ('3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4', 'How to Make Sushi', 'How to Make Sushi', 'video/how-to-make-sushi', 'thumbnail/sushi-thumbnail.jpg'),
    ('042243E2-9C87-4667-854A-AA1D6C8F64B8', 'Hiking in the Mountains', 'Hiking in the Mountains', 'video/hiking-in-the-mountains', 'thumbnail/hiking-thumbnail.jpg');

INSERT INTO tag (id, name, name_reading)
VALUES
    ('74FF4BBA-18A5-4ABC-9164-4987021D411F', 'Sushi', 'Sushi'),
    ('A082A0B6-63DE-4E97-BA30-B29319863D4F', 'Cooking', 'Cooking'),
    ('82AC868C-F677-4D08-B17C-F93B8913B69A', 'Hiking', 'Hiking'),
    ('77BABEC0-6941-4F27-9094-56F400B90B1F', 'Nature', 'Nature'),
    ('0431297E-0AB2-4E82-BE91-BB11AAB80C8B', 'Mountain', 'Mountain');

INSERT INTO video_tag (video_id, tag_id)
VALUES
    ('3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4', '74FF4BBA-18A5-4ABC-9164-4987021D411F'),
    ('3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4', 'A082A0B6-63DE-4E97-BA30-B29319863D4F'),
    ('042243E2-9C87-4667-854A-AA1D6C8F64B8', '82AC868C-F677-4D08-B17C-F93B8913B69A'),
    ('042243E2-9C87-4667-854A-AA1D6C8F64B8', '77BABEC0-6941-4F27-9094-56F400B90B1F'),
    ('042243E2-9C87-4667-854A-AA1D6C8F64B8', '0431297E-0AB2-4E82-BE91-BB11AAB80C8B');
