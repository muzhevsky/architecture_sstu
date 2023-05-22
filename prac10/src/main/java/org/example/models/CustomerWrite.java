package org.example.models;

import lombok.Getter;
import lombok.ToString;
import org.springframework.data.annotation.Id;
import org.springframework.data.relational.core.mapping.Table;

@Getter
@Table("customers")
@ToString
public class CustomerWrite {
    @Id
    private int id;
    private String surname;
    private String name;
    private String patronymic;
    private String passportNumber;
    private String passportSeries;
}