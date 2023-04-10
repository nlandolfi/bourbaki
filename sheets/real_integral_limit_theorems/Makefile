# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_integral_limit_theorems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_integral_limit_theorems.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > real_integral_limit_theorems.tex

real_integral_limit_theorems.pdf: ../../*.tex ../../trademark.pdf *.tex real_integral_limit_theorems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_integral_limit_theorems.tex
	../../latexrun real_integral_limit_theorems.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_integral_limit_theorems.pdf
	open real_integral_limit_theorems.pdf

o:
	make open

terms: real_integral_limit_theorems.tex
	bbk sheet -mode ts real_integral_limit_theorems.tex

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
	rm -f real_integral_limit_theorems.*
	make reset
	bbk sheet -mode mk > Makefile