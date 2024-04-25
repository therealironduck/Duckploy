# Use the official image as a parent image
FROM ubuntu:24.04

# Update the system, install OpenSSH Server, and set up users
RUN apt update && apt upgrade -y && \
    apt install -y openssh-server

# Create user and set password for user and root user
RUN  useradd -rm -d /home/duckploy -s /bin/bash -g root -G sudo -u 1001 duckploy && \
    echo 'duckploy:secret_password' | chpasswd && \
    echo 'root:secret_password' | chpasswd

# Set up configuration for SSH
RUN mkdir /var/run/sshd && \
    sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config && \
    sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd && \
    echo "export VISIBLE=now" >> /etc/profile

# Expose the SSH port
EXPOSE 22

# Run SSH
CMD ["/usr/sbin/sshd", "-D"]
