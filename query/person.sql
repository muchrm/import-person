-- SET SESSION group_concat_max_len = 1000000;

SELECT
			Person.personCode,
			Person.fName,
			Person.lName,
			PersonT.fName2,
			PersonT.lName2,
			PersonT.birthDate,
			rs_academic_position.acp_name_th as prefixName,
			PersonT.emailAddr AS email,
			(
				CASE Person.fStatus
				WHEN 1 THEN 'ยังปฏิบัติงานอยู่'
				WHEN 2 THEN 'ไม่ได้ปฏิบัติงานแล้ว'
				ELSE null
				END
			) AS PersonStatus,
			history_education.history_education as historyEducation,
            history_work.history_work as historyWork
		FROM Person
		LEFT JOIN PersonT ON PersonT.personId = Person.personId
		LEFT JOIN rs_academic ON rs_academic.ac_ps_id = Person.personId
		LEFT JOIN rs_academic_position ON rs_academic_position.acp_id = rs_academic.ac_acp_id
        LEFT JOIN (
            SELECT CONCAT(
            '[',
            GROUP_CONCAT(JSON_OBJECT(
			'levelName',Level.levelName,
			'levelNameEng',Level.levelNameEng,
			'degreeName',Degree.degreeName,
			'educmajorName',Educmajor.educmajorName,
			'educplaceName',Educplace.educplaceName,
			'countryName',Country.countryName
			)),
            ']') AS  history_education,
            Education.personId as personId
            FROM Education
            JOIN Level
            JOIN Degree
            JOIN Educmajor
            JOIN Educplace
            JOIN Country
            WHERE Education.levelId = Level.levelId 
            AND Education.degreeId = Degree.degreeId
            AND Education.educmajorId = Educmajor.educmajorId
            AND Education.educplaceId = Educplace.educplaceId
            AND Education.countryId = Country.countryId
            GROUP BY Education.personId
            ) AS history_education ON history_education.personId = Person.personId
		LEFT JOIN (
            SELECT CONCAT(
            '[',
            GROUP_CONCAT(JSON_OBJECT(
			'start_date',moresearcher.rs_experience.ex_start_date,
			'end_date',moresearcher.rs_experience.ex_end_date,
			'position',moresearcher.rs_experience.ex_position,
			'workplace',moresearcher.rs_experience.ex_workplace
			)),
            ']') AS  history_work,
            moresearcher.rs_user.us_ps_id AS personId
            FROM moresearcher.rs_user
			JOIN moresearcher.rs_experience
            WHERE moresearcher.rs_experience.ex_us_id = moresearcher.rs_user.us_id
            GROUP BY moresearcher.rs_user.us_ps_id
            ) AS history_work ON history_work.personId = Person.personId
