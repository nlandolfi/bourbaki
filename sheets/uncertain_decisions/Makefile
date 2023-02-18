# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make uncertain_decisions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

uncertain_decisions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > uncertain_decisions.tex

uncertain_decisions.pdf: ../../*.tex ../../trademark.pdf *.tex uncertain_decisions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode uncertain_decisions.tex
	../../latexrun uncertain_decisions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: uncertain_decisions.pdf
	open uncertain_decisions.pdf

o:
	make open

terms: uncertain_decisions.tex
	bbk_sheet -mode ts uncertain_decisions.tex

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
	rm -f uncertain_decisions.*
	make reset
	bbk_sheet -mode mk > Makefile