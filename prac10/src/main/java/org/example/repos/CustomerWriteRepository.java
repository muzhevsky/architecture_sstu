package org.example.repos;

import org.example.models.CustomerWrite;
import org.springframework.data.repository.Repository;

@org.springframework.stereotype.Repository("customerWriteRepository")
public interface CustomerWriteRepository extends Repository<CustomerWrite, Integer> {
    void save(CustomerWrite model);
}
