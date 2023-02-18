# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make monic_polynomials.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

monic_polynomials.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > monic_polynomials.tex

monic_polynomials.pdf: ../../*.tex ../../trademark.pdf *.tex monic_polynomials.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode monic_polynomials.tex
	../../latexrun monic_polynomials.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: monic_polynomials.pdf
	open monic_polynomials.pdf

o:
	make open

terms: monic_polynomials.tex
	bbk_sheet -mode ts monic_polynomials.tex

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
	rm -f monic_polynomials.*
	make reset
	bbk_sheet -mode mk > Makefile