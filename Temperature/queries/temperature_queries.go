package queries

const SaveTemperatureQuery = `
INSERT INTO TEMPERATURES (user_id, measurement, device_id) VALUES (?,?,?)
`
const GetAllTemperatures = `
SELECT measurement,date,time FROM TEMPERATURES
`
const GetByDate = `
SELECT measurement,date,time FROM TEMPERATURES
WHERE user_id = ? AND date = ?
`
const GetByUser = `
SELECT measurement,date,time FROM TEMPERATURES
WHERE user_id = ?
`
const GetLast7Days = `
SELECT measurement,date,time FROM TEMPERATURES
WHERE date >= CURDATE() - INTERVAL 7 DAY AND user_id = ? 
ORDER BY date ASC, time ASC
`
const GetForSupervisor = `
SELECT t.measurement,t.date,t.time,
u.name,u.surnames,u.email FROM TEMPERATURES t INNER JOIN USERS u
ON t.user_id = u.user_id WHERE t.user_id = ?
`