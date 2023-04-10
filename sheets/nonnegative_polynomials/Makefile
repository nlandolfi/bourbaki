# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make nonnegative_polynomials.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

nonnegative_polynomials.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > nonnegative_polynomials.tex

nonnegative_polynomials.pdf: ../../*.tex ../../trademark.pdf *.tex nonnegative_polynomials.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode nonnegative_polynomials.tex
	../../latexrun nonnegative_polynomials.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: nonnegative_polynomials.pdf
	open nonnegative_polynomials.pdf

o:
	make open

terms: nonnegative_polynomials.tex
	bbk sheet -mode ts nonnegative_polynomials.tex

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
	rm -f nonnegative_polynomials.*
	make reset
	bbk sheet -mode mk > Makefile