-- SET SESSION group_concat_max_len = 1000000;

SELECT
		    `Person`.`personId`,
            `Person`.`personCode`,
            `Person`.`fName`,
            `Person`.`lName`,
            `PersonT`.`fName2`,
            `PersonT`.`lName2`,
            `PersonT`.`birthDate`,
            `rs_academic_position`.`acp_name_th`,
            `PersonT`.`emailAddr` AS `email`,
            (
                CASE `Person`.`fStatus`
                WHEN 1 THEN 'ยังปฏิบัติงานอยู่'
                WHEN 2 THEN 'ไม่ได้ปฏิบัติงานแล้ว'
                ELSE null
                END
            ) AS `PersonStatus`,
            CONCAT(
            '[',
            GROUP_CONCAT(
            JSON_OBJECT(
            'levelName',`Level`.`levelName`,
            'levelNameEng',`Level`.`levelNameEng`,
            'degreeName',`Degree`.`degreeName`,
            'educmajorName',`Educmajor`.`educmajorName`,
            'educplaceName',`Educplace`.`educplaceName`,
            'countryName',`Country`.`countryName`
            )),']') AS `history_education`,
            CONCAT(
            '[',
            GROUP_CONCAT(
            JSON_OBJECT(
            'start_date',`moresearcher`.`rs_experience`.`ex_start_date`,
            'end_date',`moresearcher`.`rs_experience`.`ex_end_date`,
            'position',`moresearcher`.`rs_experience`.`ex_position`,
            'workplace',`moresearcher`.`rs_experience`.`ex_workplace`
            )),']') AS `history_work`
FROM `Person`
LEFT JOIN `PersonT` ON `PersonT`.`personId` = `Person`.`personId`
LEFT JOIN `rs_academic` ON `rs_academic`.`ac_ps_id` = `Person`.`personId`
LEFT JOIN `rs_academic_position` ON `rs_academic_position`.`acp_id` = `rs_academic`.`ac_acp_id`
LEFT JOIN `Education` ON `Education`.`personId` = `Person`.`personId`
LEFT JOIN `Level` ON `Education`.`levelId` = `Level`.`levelId`
LEFT JOIN `Degree` ON `Education`.`degreeId` = `Degree`.`degreeId`
LEFT JOIN `Educmajor` ON `Education`.`educmajorId` = `Educmajor`.`educmajorId`
LEFT JOIN `Educplace` ON `Education`.`educplaceId` = `Educplace`.`educplaceId`
LEFT JOIN `Country` ON `Education`.`countryId` = `Country`.`countryId`
LEFT JOIN `moresearcher`.`rs_user` ON `moresearcher`.`rs_user`.`us_ps_id` = `Person`.`personId`
LEFT JOIN `moresearcher`.`rs_experience` ON `moresearcher`.`rs_experience`.`ex_us_id` = `moresearcher`.`rs_user`.`us_id`
GROUP BY `Person`.`personId`
