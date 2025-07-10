package queries

const SaveOxygen = `
INSERT INTO OXYGENS (user_id, measurement, device_id) VALUES (?, ?, ?)
`
const GetAllOxygen = `
SELECT measurement,date,time FROM OXYGENS
`
const GetOxygenByDate = `
SELECT measurement, date, time FROM OXYGENS WHERE user_id = ? AND date = ?
`
const GetByUser = `
SELECT measurement, date, time FROM OXYGENS WHERE user_id = ?
`
const GetLastSevenDays = `
SELECT measurement, date, time FROM OXYGENS 
WHERE  date >= CURDATE() - INTERVAL 7 DAY AND user_id = ? 
ORDER BY date ASC, time ASC
`
const GetForSupervisor = `
SELECT o.measurement,o.date,o.time,
u.name,u.surnames,u.email FROM OXYGENS o INNER JOIN USERS u
ON o.user_id = u.user_id WHERE o.user_id = ?
`