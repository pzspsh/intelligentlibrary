# -*- encoding: utf-8 -*-
'''
@File   : socket5iCMPPing.py
@Time   : 2023-07-03 14:14:17
@Author : pan
'''
import socket
import os
import struct
import time
import select

# ICMP echo_request
TYPE_ECHO_REQUEST = 8
CODE_ECHO_REQUEST_DEFAULT = 0
# ICMP echo_reply
TYPE_ECHO_REPLY = 0
CODE_ECHO_REPLY_DEFAULT = 0
# ICMP overtime
TYPE_ICMP_OVERTIME = 11
CODE_TTL_OVERTIME = 0
# ICMP unreachable
TYPE_ICMP_UNREACHED = 3

MAX_HOPS = 30  # set max hops-30
TRIES = 3  # detect 3 times


# checksum
def check_sum(strings):
    csum = 0
    countTo = (len(strings) / 2) * 2
    count = 0
    while count < countTo:
        thisVal = strings[count + 1] * 256 + strings[count]
        csum = csum + thisVal
        csum = csum & 0xffffffff
        count = count + 2
    if countTo < len(strings):
        csum = csum + strings[len(strings) - 1]
        csum = csum & 0xffffffff
    csum = (csum >> 16) + (csum & 0xffff)
    csum = csum + (csum >> 16)
    answer = ~csum
    answer = answer & 0xffff
    answer = answer >> 8 | (answer << 8 & 0xff00)
    return answer


# get host_info by address
def get_host_info(host_addr):
    try:
        host_info = socket.gethostbyaddr(host_addr)
    except socket.error as e:
        display = '{0}'.format(host_addr)
    else:
        display = '{0} ({1})'.format(host_addr, host_info[0])
    return display


# construct ICMP datagram
def build_packet():
    # primitive checksum
    my_checksum = 0
    # process_id
    my_id = os.getpid()
    # sequence as 1(>0)
    my_seq = 1
    # 2's header
    my_header = struct.pack("bbHHh", TYPE_ECHO_REQUEST, CODE_ECHO_REQUEST_DEFAULT, my_checksum, my_id, my_seq)
    # SYS_time as payload
    my_data = struct.pack("d", time.time())
    # temporary datagram
    package = my_header + my_data
    # true checksum
    my_checksum = check_sum(package)
    # windows-big endian
    my_checksum = socket.htons(my_checksum)
    # repack
    my_header = struct.pack("bbHHh", TYPE_ECHO_REQUEST, CODE_ECHO_REQUEST_DEFAULT, my_checksum, my_id, 1)
    # true datagram
    ip_package = my_header + my_data
    return ip_package


def main(hostname):
    print("routing {0}[{1}](max hops = 30, detect tries = 3)".format(hostname, socket.gethostbyname(hostname)))
    for ttl in range(1, MAX_HOPS):
        print("%2d" % ttl, end="")
        for tries in range(0, TRIES):
            # create raw socket
            icmp_socket = socket.socket(socket.AF_INET, socket.SOCK_RAW, socket.getprotobyname("icmp"))
            icmp_socket.setsockopt(socket.IPPROTO_IP, socket.IP_TTL, struct.pack('I', ttl))
            icmp_socket.settimeout(TIMEOUT)
            # construct datagram
            icmp_package = build_packet()
            icmp_socket.sendto(icmp_package, (hostname, 0))
            # waiting for receiving reply
            start_time = time.time()
            select.select([icmp_socket], [], [], TIMEOUT)
            end_time = time.time()
            # compute time of receiving
            during_time = end_time - start_time
            if during_time >= TIMEOUT or during_time == 0:
                print("    *    ", end="")
            else:
                print(" %4.0f ms " % (during_time * 1000), end="")
            if tries >= TRIES - 1:
                try:
                    ip_package, ip_info = icmp_socket.recvfrom(1024)
                except socket.timeout:
                    print(" request time out")
                else:
                    # extract ICMP header from IP datagram
                    icmp_header = ip_package[20:28]

                    # unpack ICMP header
                    after_type, after_code, after_checksum, after_id, after_sequence = struct.unpack("bbHHh", icmp_header)
                    output = get_host_info(ip_info[0])

                    if after_type == TYPE_ICMP_UNREACHED:  # unreachable
                        print("Wrong!unreached net/host/port!")
                        break
                    elif after_type == TYPE_ICMP_OVERTIME:  # ttl overtimr
                        print(" %s" % output)
                        continue
                    elif after_type == 0:  # type_echo
                        print(" %s" % output)
                        print("program run over!")
                        return
                    else:
                        print("request timeout")
                        print("program run wrongly!")
                        return


if __name__ == "__main__":
    while True:
        try:
            ip = input("please input a ip address:")
            global TIMEOUT
            TIMEOUT = int(input("Input timeout you want: "))
            main(ip)
            break
        except Exception as e:
            print(e)
            continue