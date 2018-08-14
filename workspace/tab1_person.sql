SELECT
                                        Person.*,
                                        Prefix.prefixName,
                                        Prefix.prefixNameEng,
                                        Prefix.defaultSex,
                                        IF(Prefix.defaultSex='', 'ไม่ระบุ', (IF(Prefix.defaultSex='M','ชาย','หญิง'))) AS sexName
                                FROM
                                        Person
                                Left Join Prefix ON Person.prefixId = Prefix.prefixId
                                WHERE
                                        Person.personId = 379
