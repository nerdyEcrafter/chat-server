package org.wid.userservice.adapter.primary.rest;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.wid.userservice.dto.user.RegisterUserDto;
import org.wid.userservice.dto.user.UserDto;
import org.wid.userservice.port.primary.UserServicePort;

import lombok.extern.slf4j.Slf4j;

@RestController
@RequestMapping("/api/v1")
@Slf4j
public class UserController {
  private final UserServicePort userService;

  public UserController(UserServicePort userService) {
    this.userService = userService;
  }

  @PostMapping("/user")
  public ResponseEntity<String> register(@Validated @RequestBody RegisterUserDto registerUserDto) {
    userService.register(registerUserDto);

    return ResponseEntity.status(HttpStatus.NOT_IMPLEMENTED).body("register");
  }

  public ResponseEntity<String> login() {

    return ResponseEntity.status(HttpStatus.NOT_IMPLEMENTED).body("login");
  }

  @GetMapping("/user/{userId}")
  public ResponseEntity<UserDto> getUser(@PathVariable long userId) {
    log.info("get user id:{}", userId);
    UserDto userDto = userService.getUser(userId);

    return ResponseEntity.ok(userDto);
  }
}
