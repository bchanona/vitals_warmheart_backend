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