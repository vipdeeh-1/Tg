//
// Created by Laky64 on 23/04/2025.
// This header file provides compatibility for the glibc >= 2.28 version for __dn_expand and __res_nquery functions.

#pragma once

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __GLIBC__
    #if __GLIBC__ > 2 || (__GLIBC__ == 2 && __GLIBC_MINOR__ >= 28)
		#include <resolv.h>

		__attribute__((weak))
		int __dn_expand(const unsigned char *msg, const unsigned char *eomorig,
		                 const unsigned char *comp_dn, char *exp_dn, int length) {
		    return dn_expand(msg, eomorig, comp_dn, exp_dn, length);
		}

		__attribute__((weak))
		int __res_nquery(res_state statp, const char *dname, int class, int type,
		                 unsigned char *answer, int anslen) {
		    return res_nquery(statp, dname, class, type, answer, anslen);
		}
	#endif
#endif

#ifdef __cplusplus
}
#endif
