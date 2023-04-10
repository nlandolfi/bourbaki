# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make derivatives_of_scalar_multiples.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

derivatives_of_scalar_multiples.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > derivatives_of_scalar_multiples.tex

derivatives_of_scalar_multiples.pdf: ../../*.tex ../../trademark.pdf *.tex derivatives_of_scalar_multiples.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode derivatives_of_scalar_multiples.tex
	../../latexrun derivatives_of_scalar_multiples.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: derivatives_of_scalar_multiples.pdf
	open derivatives_of_scalar_multiples.pdf

o:
	make open

terms: derivatives_of_scalar_multiples.tex
	bbk sheet -mode ts derivatives_of_scalar_multiples.tex

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
	rm -f derivatives_of_scalar_multiples.*
	make reset
	bbk sheet -mode mk > Makefile