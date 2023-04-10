# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make ordered_pairs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

ordered_pairs.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > ordered_pairs.tex

ordered_pairs.pdf: ../../*.tex ../../trademark.pdf *.tex ordered_pairs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode ordered_pairs.tex
	../../latexrun ordered_pairs.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: ordered_pairs.pdf
	open ordered_pairs.pdf

o:
	make open

terms: ordered_pairs.tex
	bbk sheet -mode ts ordered_pairs.tex

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
	rm -f ordered_pairs.*
	make reset
	bbk sheet -mode mk > Makefile