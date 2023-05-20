package org.example.controllers;

import com.google.gson.Gson;
import org.example.repos.CustomerReadRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class CustomerReadController {
    @Autowired
    @Qualifier("gson")
    Gson gson;

    @Autowired
    @Qualifier("customerReadRepository")
    CustomerReadRepository repository;

    @GetMapping("/customer/{id}")
    public ResponseEntity<String> getCustomer(@PathVariable int id){
        var model = repository.findById(id);
        return new ResponseEntity<>(gson.toJson(model), HttpStatus.OK);
    }
}