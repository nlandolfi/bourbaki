# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_polynomials.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_polynomials.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > real_polynomials.tex

real_polynomials.pdf: ../../*.tex ../../trademark.pdf *.tex real_polynomials.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_polynomials.tex
	../../latexrun real_polynomials.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_polynomials.pdf
	open real_polynomials.pdf

o:
	make open

terms: real_polynomials.tex
	bbk sheet -mode ts real_polynomials.tex

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
	rm -f real_polynomials.*
	make reset
	bbk sheet -mode mk > Makefile