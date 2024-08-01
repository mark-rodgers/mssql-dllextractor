SELECT
  db.name AS [Name],
  sp.name AS [Owner]
FROM  sys.databases AS db
LEFT JOIN sys.server_principals AS sp ON db.owner_sid = sp.sid
WHERE db.name NOT IN (%s)
ORDER BY  db.name;