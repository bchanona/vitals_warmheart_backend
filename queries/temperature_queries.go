package queries

const SaveTemperatureQuery = `
INSERT INTO TEMPERATURES (user_id, measurement, device_id) VALUES (?,?,?)
`
const GetAllTemperatures = `
SELECT * FROM TEMPERATURES
`
const GetByDate = `
SELECT measuremnt,date,time FROM TEMPERATURES
WHERE user_id = ? AND date = ?
`
const GetByUser = `
SELECT measuremnt,date,time FROM TEMPERATURES
WHERE user_id = ?
`
const GetForSupervisor = `
SELECT t.measurement,t.date,t.time,
u.name,u.surnames,u.email FROM TEMPERATURES t INNER JOIN USERS u
ON t.user_id = u.user_id WHERE t.user_id = ?
`