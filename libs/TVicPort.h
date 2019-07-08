// ========================================================================
// =================    TVicPort  DLL interface        ====================
// ==========                 Version  4.0                      ===========
// ==========     Copyright (c) 1997-2005, EnTech Taiwan        ===========
// ========================================================================
// ==========           http://www.entechtaiwan.com             ===========
// ==========          mailto: tools@entechtaiwan.com           ===========
// ========================================================================


#ifndef __TVicPort_H
#define __TVicPort_H

#define VICFN  __stdcall

#ifdef __cplusplus
extern "C" {
#endif

#pragma pack(1)

typedef struct _HDDInfo {
	unsigned long	BufferSize;
	unsigned long	DoubleTransfer;
	unsigned long	ControllerType;
	unsigned long	ECCMode;
	unsigned long	SectorsPerInterrupt;
	unsigned long	Cylinders;
	unsigned long	Heads;
	unsigned long	SectorsPerTrack;
	char	Model[41];
	char	SerialNumber[21];
	char	Revision[9];
} HDDInfo, *pHDDInfo;



void	VICFN CloseTVicPort();
int	VICFN OpenTVicPort();
int	VICFN IsDriverOpened();

int	VICFN TestHardAccess();
void	VICFN SetHardAccess(int bNewValue);

unsigned char	VICFN ReadPort  (unsigned short PortAddr);
void	VICFN WritePort (unsigned short PortAddr, unsigned char nNewValue);
unsigned short	VICFN ReadPortW (unsigned short PortAddr);
void	VICFN WritePortW(unsigned short PortAddr, unsigned short nNewValue);
unsigned long	VICFN ReadPortL (unsigned short PortAddr);
void	VICFN WritePortL(unsigned short PortAddr, unsigned long nNewValue);

void	VICFN ReadPortFIFO  (unsigned short PortAddr, unsigned long NumPorts, unsigned char  * Buffer);
void	VICFN WritePortFIFO (unsigned short PortAddr, unsigned long NumPorts, unsigned char  * Buffer);
void	VICFN ReadPortWFIFO (unsigned short PortAddr, unsigned long NumPorts, unsigned short * Buffer);
void	VICFN WritePortWFIFO(unsigned short PortAddr, unsigned long NumPorts, unsigned short * Buffer);
void	VICFN ReadPortLFIFO (unsigned short PortAddr, unsigned long NumPorts, unsigned long  * Buffer);
void	VICFN WritePortLFIFO(unsigned short PortAddr, unsigned long NumPorts, unsigned long  * Buffer);

unsigned short	VICFN GetLPTNumber();
void	VICFN SetLPTNumber(unsigned short nNewValue);
unsigned short	VICFN GetLPTNumPorts();
unsigned short	VICFN GetLPTBasePort();
unsigned char	VICFN AddNewLPT(unsigned short PortBaseAddress);

int	VICFN GetPin(unsigned short nPin);
void	VICFN SetPin(unsigned short nPin, int bNewValue);

int	VICFN GetLPTAckwl();
int	VICFN GetLPTBusy();
int	VICFN GetLPTPaperEnd();
int	VICFN GetLPTSlct();
int	VICFN GetLPTError();

void	VICFN LPTInit();
void	VICFN LPTSlctIn();
void	VICFN LPTStrobe();
void	VICFN LPTAutofd(int Flag);

void	VICFN GetHDDInfo  (unsigned char IdeNumber, unsigned char Master, pHDDInfo Info);

unsigned long	VICFN MapPhysToLinear(unsigned long PhAddr, unsigned long PhSize);
void	VICFN UnmapMemory    (unsigned long PhAddr, unsigned long PhSize);

unsigned char	VICFN GetMem (unsigned long MappedAddr, unsigned long Offset);
void	VICFN SetMem (unsigned long MappedAddr, unsigned long Offset, unsigned char nNewValue);
unsigned short	VICFN GetMemW(unsigned long MappedAddr, unsigned long Offset);
void	VICFN SetMemW(unsigned long MappedAddr, unsigned long Offset, unsigned short nNewValue);
unsigned long	VICFN GetMemL(unsigned long MappedAddr, unsigned long Offset);
void	VICFN SetMemL(unsigned long MappedAddr, unsigned long Offset, unsigned long nNewValue);

void	VICFN SetLPTReadMode();
void	VICFN SetLPTWriteMode();

void	VICFN LaunchWeb();
void	VICFN LaunchMail();

int VICFN EvaluationDaysLeft();

#pragma pack()

#ifdef __cplusplus
} // extern "C"
#endif

#endif
