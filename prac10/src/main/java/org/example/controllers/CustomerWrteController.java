package org.example.controllers;

import org.example.models.CustomerWrite;
import org.example.repos.CustomerWriteRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class CustomerWrteController {
    @Autowired
    @Qualifier("customerWriteRepository")
    CustomerWriteRepository repository;

    @PostMapping("/customer")
    public ResponseEntity createCustomer(@RequestBody CustomerWrite model){
        repository.save(model);
        return new ResponseEntity(HttpStatus.OK);
    }
}
