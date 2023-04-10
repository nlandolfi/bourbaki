# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_vector_bases.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_vector_bases.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > real_vector_bases.tex

real_vector_bases.pdf: ../../*.tex ../../trademark.pdf *.tex real_vector_bases.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_vector_bases.tex
	../../latexrun real_vector_bases.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_vector_bases.pdf
	open real_vector_bases.pdf

o:
	make open

terms: real_vector_bases.tex
	bbk_sheet -mode ts real_vector_bases.tex

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
	rm -f real_vector_bases.*
	make reset
	bbk_sheet -mode mk > Makefile