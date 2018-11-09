-- Source: https://leetcode.com/problems/not-boring-movies

-- Table: cinema
-- | id | movie      | description | rating |
-- | -- | ---------- | ----------- | ------ |
-- | 1  | War        | great 3D    | 8.9    |
-- | 2  | Science    | fiction     | 8.5    |
-- | 3  | irish      | boring      | 6.2    |
-- | 4  | Ice song   | Fantacy     | 8.6    |
-- | 5  | House card | Interesting | 9.1    |

SELECT `id`, `movie`, `description`, `rating`
FROM `cinema`
WHERE `id`%2 = 1 AND `description` != 'boring'
ORDER BY `rating` DESC;
