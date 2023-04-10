# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make ordered_pair_pathologies.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

ordered_pair_pathologies.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > ordered_pair_pathologies.tex

ordered_pair_pathologies.pdf: ../../*.tex ../../trademark.pdf *.tex ordered_pair_pathologies.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode ordered_pair_pathologies.tex
	../../latexrun ordered_pair_pathologies.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: ordered_pair_pathologies.pdf
	open ordered_pair_pathologies.pdf

o:
	make open

terms: ordered_pair_pathologies.tex
	bbk sheet -mode ts ordered_pair_pathologies.tex

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
	rm -f ordered_pair_pathologies.*
	make reset
	bbk sheet -mode mk > Makefile