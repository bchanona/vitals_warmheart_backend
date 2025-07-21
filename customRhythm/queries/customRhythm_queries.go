package queries

const SaveCustomQuery = `
INSERT INTO HEART_RATE_PERSONALIZED (user_id,bpm_low,bpm_high) VALUES (?,?,?)
`
const UpdateCustom =`
UPDATE HEART_RATE_PERSONALIZED
SET bpm_low = ?, bpm_high = ? 
WHERE user_id = ?
`