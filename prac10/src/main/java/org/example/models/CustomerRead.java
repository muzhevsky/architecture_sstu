package org.example.models;

import lombok.ToString;
import org.springframework.data.annotation.Id;
import org.springframework.data.relational.core.mapping.Table;

@Table("customers_view")
@ToString
public class CustomerRead {
    @Id
    private int id;
    private String fullname;
    private String passport;
}
