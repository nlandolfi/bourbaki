# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make ordinary_reducer_factorization.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

ordinary_reducer_factorization.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > ordinary_reducer_factorization.tex

ordinary_reducer_factorization.pdf: ../../*.tex ../../trademark.pdf *.tex ordinary_reducer_factorization.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode ordinary_reducer_factorization.tex
	../../latexrun ordinary_reducer_factorization.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: ordinary_reducer_factorization.pdf
	open ordinary_reducer_factorization.pdf

o:
	make open

terms: ordinary_reducer_factorization.tex
	bbk sheet -mode ts ordinary_reducer_factorization.tex

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
	rm -f ordinary_reducer_factorization.*
	make reset
	bbk sheet -mode mk > Makefile