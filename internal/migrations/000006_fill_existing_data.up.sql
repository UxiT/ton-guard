insert into profile values (
                            '55e0ce1f-a524-4e97-8735-7632ce343883',
                            243244613,
                            'test@gmail.com',
                            '$2a$10$Fzwn9vTYfUFLHbnu15ZIvuZiEWpSSp.hc9NgrT/ZEsxtV1wvzhg9G',
                        now(),
                            now()
                           );

insert into customer values (
                             '477325b4-6ec9-46ee-98d9-f87c6a2adbb5',
                             '55e0ce1f-a524-4e97-8735-7632ce343883',
                             now(),
                             now()
                            );

insert into account values (
                            '924f9548-f912-49ff-85eb-6e34dcccac63',
                            '321a0ccf-ec45-462b-94e7-cc4be892348d',
                            'eur',
                            'active',
                            '0.000000000000000000',
                            '477325b4-6ec9-46ee-98d9-f87c6a2adbb5',
                            now(),
                            now()
                           )