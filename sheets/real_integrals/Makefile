# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_integrals.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_integrals.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > real_integrals.tex

real_integrals.pdf: ../../*.tex ../../trademark.pdf *.tex real_integrals.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_integrals.tex
	../../latexrun real_integrals.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_integrals.pdf
	open real_integrals.pdf

o:
	make open

terms: real_integrals.tex
	bbk_sheet -mode ts real_integrals.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f real_integrals.*
	make reset
	bbk_sheet -mode mk > Makefile