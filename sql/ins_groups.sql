INSERT OR IGNORE INTO Groups (
    structureId,
    facultyId,
    id,
    name,
    course,
    priority,
    educationForm
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
);
