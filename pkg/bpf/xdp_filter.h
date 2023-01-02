/*
*	Copyright (C) 2023 Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

#include <net/ethernet.h>

#ifndef BPF_LIBRARY_H
#define BPF_LIBRARY_H

#endif //BPF_LIBRARY_H

#define MAC_BYTES = 6;
#define IPV4_BYTES = 32;
#define IPV6_BYTES = 128;

enum STACK_PUSH {
    FORWARD_UP = 0,
    FORWARD_DROP,
    FORWARD_OUT,
    FORWARD_RETURN
};

struct ipv4_hdr {

};

struct ipv6_hdr {

};

struct tcp_hdr {

};

struct udp_hdr {

};