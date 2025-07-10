package queries

const SaveHeartRateQuery = `
INSERT INTO HEART_RATES (user_id, measurement, device_id) VALUES (?,?,?)
`
const GetAllHeartRates = `
SELECT measurement,date,time FROM HEART_RATES
`
const GetByDate = `
SELECT measurement,date,time FROM HEART_RATES
WHERE user_id = ? AND date = ?
`
const GetByUser = `
SELECT measurement,date,time FROM HEART_RATES
WHERE user_id = ?
`
const GetLast7Days = `
SELECT measurement,date,time FROM HEART_RATES
WHERE date >= CURDATE() - INTERVAL 7 DAY AND user_id = ? 
ORDER BY date ASC, time ASC
`
const GetForSupervisor = `
SELECT h.measurement,h.date,h.time,
u.name,u.surnames,u.email FROM HEART_RATES h INNER JOIN USERS u
ON h.user_id = u.user_id WHERE h.user_id = ?
`