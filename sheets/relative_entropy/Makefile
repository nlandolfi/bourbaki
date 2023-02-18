# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make relative_entropy.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

relative_entropy.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > relative_entropy.tex

relative_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex relative_entropy.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode relative_entropy.tex
	../../latexrun relative_entropy.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: relative_entropy.pdf
	open relative_entropy.pdf

o:
	make open

terms: relative_entropy.tex
	bbk_sheet -mode ts relative_entropy.tex

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
	rm -f relative_entropy.*
	make reset
	bbk_sheet -mode mk > Makefile