# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_modular_arithmetic.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_modular_arithmetic.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > real_modular_arithmetic.tex

real_modular_arithmetic.pdf: ../../*.tex ../../trademark.pdf *.tex real_modular_arithmetic.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_modular_arithmetic.tex
	../../latexrun real_modular_arithmetic.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_modular_arithmetic.pdf
	open real_modular_arithmetic.pdf

o:
	make open

terms: real_modular_arithmetic.tex
	bbk sheet -mode ts real_modular_arithmetic.tex

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
	rm -f real_modular_arithmetic.*
	make reset
	bbk sheet -mode mk > Makefile