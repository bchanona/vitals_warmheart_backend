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
