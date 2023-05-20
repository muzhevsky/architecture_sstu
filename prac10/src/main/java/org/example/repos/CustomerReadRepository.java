package org.example.repos;

import org.example.models.CustomerRead;
import org.springframework.data.repository.Repository;

@org.springframework.stereotype.Repository("customerReadRepository")
public interface CustomerReadRepository extends Repository<CustomerRead, Integer> {
    CustomerRead findById(Integer id);
}
