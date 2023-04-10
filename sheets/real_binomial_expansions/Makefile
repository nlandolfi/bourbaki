# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_binomial_expansions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_binomial_expansions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > real_binomial_expansions.tex

real_binomial_expansions.pdf: ../../*.tex ../../trademark.pdf *.tex real_binomial_expansions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_binomial_expansions.tex
	../../latexrun real_binomial_expansions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_binomial_expansions.pdf
	open real_binomial_expansions.pdf

o:
	make open

terms: real_binomial_expansions.tex
	bbk sheet -mode ts real_binomial_expansions.tex

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
	rm -f real_binomial_expansions.*
	make reset
	bbk sheet -mode mk > Makefile