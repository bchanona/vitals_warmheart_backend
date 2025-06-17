package queries

const SaveTemperatureQuery = `
INSERT INTO TEMPERATURES (user_id, measurement, device_id) VALUES (?,?,?)
`